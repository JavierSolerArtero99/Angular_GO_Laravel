import { Component, Input, OnInit } from "@angular/core";
import { Comment } from "../../core";

@Component({
  selector: "app-product-comment",
  templateUrl: "./product-comment.component.html",
  styles: [],
})
export class ProductCommentComponent implements OnInit {
  @Input() comment: Comment;
  constructor() {}

  ngOnInit(): void {
  }
}
