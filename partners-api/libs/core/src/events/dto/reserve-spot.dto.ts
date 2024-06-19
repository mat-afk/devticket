import { TicketKind } from '@prisma/client';

export class ReserveSpotDto {
  spots: string[];
  email: string;
  ticketKind: TicketKind;
}
