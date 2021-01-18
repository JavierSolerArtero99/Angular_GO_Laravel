import { Comment } from "../../core";

export interface Product {
  Id: number;
  name: string;
  likes: number;
  comments: Comment[];
  //   user: number
}