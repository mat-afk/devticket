import { Module } from '@nestjs/common';
import { PrismaModule } from '@app/core/prisma/prisma.module';
import { ConfigModule } from '@nestjs/config';
import { EventosModule } from './eventos/eventos.module';
import { LugaresModule } from './lugares/lugares.module';
import { AuthModule } from '@app/core/auth/auth.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env.partner2',
      isGlobal: true,
    }),
    PrismaModule,
    EventosModule,
    LugaresModule,
    AuthModule,
  ],
})
export class Partner2Module {}
