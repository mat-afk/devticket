import { Module } from '@nestjs/common';
import { PrismaModule } from '@app/core/prisma/prisma.module';
import { SpotsModule } from './spots/spots.module';
import { ConfigModule } from '@nestjs/config';
import { EventsModule } from './events/events.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env.partner1',
    }),
    PrismaModule,
    EventsModule,
    SpotsModule,
  ],
})
export class Partner1Module {}
