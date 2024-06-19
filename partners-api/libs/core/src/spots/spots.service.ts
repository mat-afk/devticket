import { Injectable } from '@nestjs/common';
import { CreateSpotDto } from './dto/create-spot.dto';
import { UpdateSpotDto } from './dto/update-spot.dto';
import { PrismaService } from '../prisma/prisma.service';

@Injectable()
export class SpotsService {
  constructor(private prismaService: PrismaService) {}

  async create(createSpotDto: CreateSpotDto & { eventId: string }) {
    const event = await this.prismaService.event.findFirst({
      where: { id: createSpotDto.eventId },
    });

    if (!event) {
      throw new Error('Event not found.');
    }

    return this.prismaService.spot.create({
      data: createSpotDto,
    });
  }

  findAll(eventId: string) {
    return this.prismaService.spot.findMany({
      where: { eventId },
    });
  }

  findOne(spotId: string, eventId: string) {
    return this.prismaService.spot.findUnique({
      where: {
        id: spotId,
        eventId,
      },
    });
  }

  update(spotId: string, eventId: string, updateSpotDto: UpdateSpotDto) {
    return this.prismaService.spot.update({
      data: updateSpotDto,
      where: {
        id: spotId,
        eventId,
      },
    });
  }

  remove(spotId: string, eventId: string) {
    return this.prismaService.spot.delete({
      where: {
        id: spotId,
        eventId,
      },
    });
  }
}
