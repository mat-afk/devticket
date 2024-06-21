import { TicketKind } from '@prisma/client';

export class ReserveSpotRequest {
  spots: string[];
  email: string;
  ticketKind: TicketKind;
}
