import { TicketKind } from '@prisma/client';

export class ReservarLugarRequest {
  lugares: string[];
  email: string;
  tipoIngresso: TicketKind;
}
