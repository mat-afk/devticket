import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { EventsModule } from '../../partner1/src/events/events.module';
import { PrismaModule } from './prisma/prisma.module';
import { SpotsModule } from '../../partner1/src/spots/spots.module';

@Module({
  imports: [EventsModule, PrismaModule, SpotsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
