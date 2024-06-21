import { PartialType } from '@nestjs/mapped-types';
import { CreateSpotRequest } from './create-spot.request';
import { Status } from '@prisma/client';

export class UpdateSpotRequest extends PartialType(CreateSpotRequest) {
  name: string;
  status?: Status;
}
