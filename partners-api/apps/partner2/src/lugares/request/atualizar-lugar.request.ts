import { PartialType } from '@nestjs/mapped-types';
import { CriarLugarRequest } from './criar-lugar.request';
import { Status } from '@prisma/client';

export class AtualizarLugarRequest extends PartialType(CriarLugarRequest) {
  nome: string;
  estado?: Status;
}
