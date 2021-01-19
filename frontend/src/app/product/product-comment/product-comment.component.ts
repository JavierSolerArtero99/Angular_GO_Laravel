import { Component, Input, OnInit } from "@angular/core";
import { Comment } from "../../core";
import { ProductService } from "../shared/product.service";

@Component({
  selector: "app-product-comment",
  templateUrl: "./product-comment.component.html",
  styles: [],
})
export class ProductCommentComponent implements OnInit {
  @Input() comment: Comment;
  @Input() authorId: any;
  @Input() deleteComment: any;
  constructor(private productService: ProductService) {}

  ngOnInit(): void {}

  isAuthor(): boolean {
    return this.comment.Author.ID === this.authorId;
  }

  removeComment(): void {
    console.log(this.comment.Author.ID, this.comment.Message);

    this.deleteComment(this.comment);
  }
}
