import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  HttpCode,
} from '@nestjs/common';
import { SpotsService } from '@app/core/spots/spots.service';
import { CriarLugarRequest } from './request/criar-lugar.request';
import { AtualizarLugarRequest } from './request/atualizar-lugar.request';
import { Status } from '@prisma/client';

@Controller('eventos/:eventoId/lugares')
export class LugaresController {
  constructor(private readonly spotsService: SpotsService) {}

  @Post()
  create(
    @Param('eventoId') eventoId: string,
    @Body() criarLugarRequest: CriarLugarRequest,
  ) {
    return this.spotsService.create({
      eventId: eventoId,
      name: criarLugarRequest.nome,
    });
  }

  @Get()
  findAll(@Param('eventoId') eventoId: string) {
    return this.spotsService.findAll(eventoId);
  }

  @Get(':lugarId')
  findOne(
    @Param('lugarId') lugarId: string,
    @Param('eventoId') eventoId: string,
  ) {
    return this.spotsService.findOne(lugarId, eventoId);
  }

  @Patch(':lugarId')
  update(
    @Param('lugarId') lugarId: string,
    @Param('eventoId') eventoId: string,
    @Body() atualizarLugarRequest: AtualizarLugarRequest,
  ) {
    return this.spotsService.update(lugarId, eventoId, {
      name: atualizarLugarRequest.nome,
      status:
        atualizarLugarRequest.estado === 'disponível'
          ? Status.available
          : Status.reserved,
    });
  }

  @HttpCode(204)
  @Delete(':lugarId')
  remove(
    @Param('lugarId') lugarId: string,
    @Param('eventoId') eventoId: string,
  ) {
    return this.spotsService.remove(lugarId, eventoId);
  }
}
