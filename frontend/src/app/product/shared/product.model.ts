import { Comment } from "../../core";

export interface Product {
  Id: number;
  Name: string;
  image: string;
  Price: number;
  Likes: number;
  LikesList: any;
  Comments: Comment[];
  //   user: number
}