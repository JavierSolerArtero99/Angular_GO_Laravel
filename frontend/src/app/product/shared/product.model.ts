import { Comment } from "../../core";

export interface Product {
  Id: number;
  Name: string;
  image: string;
  Price: number;
  likes: number;
  Comments: Comment[];
  //   user: number
}