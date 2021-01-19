import { Profile } from './profile.model';

export interface Comment {
  id: number;
  Message: string;
  createdAt: string;
  Author: Profile;
}
