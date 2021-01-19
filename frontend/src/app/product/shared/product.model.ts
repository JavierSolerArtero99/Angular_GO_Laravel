import { Comment } from "../../core";

export interface Product {
  Id: number;
  name: string;
  image: string;
  price: number;
  likes: number;
  Comments: Comment[];
  //   user: number
}