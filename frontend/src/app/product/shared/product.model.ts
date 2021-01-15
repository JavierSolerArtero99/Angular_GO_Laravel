import { Comment } from "../../core";

export interface Product {
  id: number;
  name: string;
  likes: number;
  comments: Comment[];
  //   user: number
}