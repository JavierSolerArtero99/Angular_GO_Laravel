import { Component, Input, OnInit } from "@angular/core";
import { Comment } from "../../core";

@Component({
  selector: "app-product-comment",
  templateUrl: "./product-comment.component.html",
  styles: [],
})
export class ProductCommentComponent implements OnInit {
  @Input() comment: Comment;
  @Input() authorId: any;
  @Input() deleteComment: any;
  constructor() {}

  ngOnInit(): void {}

  isAuthor(): boolean {
    return this.comment.Author.ID === this.authorId;
  }

  removeComment(): void {
    this.deleteComment(this.comment)
  }
}
