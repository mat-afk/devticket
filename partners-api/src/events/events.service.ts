import { Injectable } from '@nestjs/common';
import { CreateEventDto } from './dto/create-event.dto';
import { UpdateEventDto } from './dto/update-event.dto';
import { PrismaService } from 'src/prisma/prisma.service';
import { ReserveSpotDto } from './dto/reserve-spot.dto';
import { Prisma, Status, TicketStatus } from '@prisma/client';

@Injectable()
export class EventsService {
  constructor(private prismaService: PrismaService) {}

  create(createEventDto: CreateEventDto) {
    return this.prismaService.event.create({
      data: {
        ...createEventDto,
        date: new Date(createEventDto.date),
      },
    });
  }

  findAll() {
    return this.prismaService.event.findMany();
  }

  findOne(id: string) {
    return this.prismaService.event.findUnique({
      where: { id },
    });
  }

  update(id: string, updateEventDto: UpdateEventDto) {
    return this.prismaService.event.update({
      data: {
        ...updateEventDto,
        date: new Date(updateEventDto.date),
      },
      where: { id },
    });
  }

  remove(id: string) {
    return this.prismaService.event.delete({
      where: { id },
    });
  }

  async reserveSpot(reserveSpotDto: ReserveSpotDto & { eventId: string }) {
    const spots = await this.prismaService.spot.findMany({
      where: {
        eventId: reserveSpotDto.eventId,
        name: {
          in: reserveSpotDto.spots,
        },
      },
    });

    if (spots.length !== reserveSpotDto.spots.length) {
      const foundSpots = spots.map((spot) => spot.name);
      const notFoundSpots = reserveSpotDto.spots.filter(
        (spotName) => !foundSpots.includes(spotName),
      );

      throw new Error(`Spots ${notFoundSpots.join(', ')} were not found`);
    }

    try {
      return await this.prismaService.$transaction(
        async (prisma) => {
          await prisma.reservationHistory.createMany({
            data: spots.map((spot) => ({
              spotId: spot.id,
              email: reserveSpotDto.email,
              ticketKind: reserveSpotDto.ticketKind,
              status: TicketStatus.reserved,
            })),
          });

          await prisma.spot.updateMany({
            data: {
              status: Status.reserved,
            },
            where: {
              id: {
                in: spots.map((spot) => spot.id),
              },
            },
          });

          const tickets = Promise.all(
            spots.map((spot) =>
              prisma.ticket.create({
                data: {
                  email: reserveSpotDto.email,
                  ticketKind: reserveSpotDto.ticketKind,
                  spotId: spot.id,
                },
              }),
            ),
          );
          return tickets;
        },
        { isolationLevel: Prisma.TransactionIsolationLevel.ReadCommitted },
      );
    } catch (e) {
      if (e instanceof Prisma.PrismaClientKnownRequestError) {
        switch (e.code) {
          case 'P2002': // Unique constraint violation
          case 'P2034': // Transaction conflict
            throw new Error('Some spots are already reserved.');
        }
      }
      throw e;
    }
  }
}
