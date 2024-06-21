import { Module } from '@nestjs/common';
import { PrismaModule } from '@app/core/prisma/prisma.module';
import { SpotsModule } from './spots/spots.module';
import { ConfigModule } from '@nestjs/config';
import { EventsModule } from './events/events.module';
import { AuthModule } from '@app/core/auth/auth.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env.partner1',
      isGlobal: true,
    }),
    PrismaModule,
    EventsModule,
    SpotsModule,
    AuthModule,
  ],
})
export class Partner1Module {}
