import { Component, OnInit } from "@angular/core";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
import { JwtService, User, UserService } from "../../core";
import { Product } from "../shared/product.model";
import { ProductService } from "../shared/product.service";
import { Comment } from "../../core";

@Component({
  selector: "app-product-details",
  templateUrl: "./product-details.component.html",
})
export class ProductDetailsComponent implements OnInit {
  commentErrorMessage: String;
  commentSuccessMessage: String;
  product: Product;
  currentUser: User;
  jwtService: JwtService;
  canModify: boolean;
  comments: Comment[];
  form: FormGroup;
  commentMessage = new FormControl();
  commentFormErrors = {};
  isSubmitting = false;
  isDeleting = false;
  loading: boolean = true;
  showBuy: boolean = false;
  showError: boolean = false;

  constructor(
    private productService: ProductService,
    private fb: FormBuilder,
    private userService: UserService,
    // private articlesService: ArticlesService,
    // private commentsService: CommentsService,
    private router: Router // private userService: UserService
  ) {
    this.jwtService = new JwtService();
  }

  ngOnInit() {
    // slicepara cojer el nombre del producto
    let productName = this.router.url.substr(9);

    this.currentUser = this.userService.getCurrentUser();
    

    // obteniendo los datos del producto
    this.productService.getSingleProduct(productName).subscribe((data) => {
      this.product = data.product;
      console.log(this.product);

      this.loading = false;
      this.comments = data.product.Comments;
      console.log("===data===");
      console.log(this.comments);
    });

    this.form = this.fb.group({
      commentMessage: "",
    });
  }

  publishComment() {
    let token = this.jwtService.getToken();

    if (!token) {
      this.commentErrorMessage =
        "Tienes que iniciar sesión para publicar un comentario";
      return;
    }

    let comment = {
      UserID: this.userService.getCurrentUser().id,
      ProductID: this.product.Id,
      Message: this.form.getRawValue().commentMessage,
    };

    if (comment.Message.length <= 0) {
      this.commentErrorMessage = "Introduce un comentario";
      return;
    }

    this.productService.postComment(comment).subscribe((data) => {
      this.commentErrorMessage = "";
      this.commentSuccessMessage = "Comentario publicado";
    });
  }

  getAuthorId(): number {
    return this.userService.getCurrentUser().id;
  }

  buyProduct() {
    console.log(this.product);

    if (!this.currentUser.token && !this.userService.getCurrentUser()) {this.showError = true; return };
    
    // comprando producto
    this.productService
      .buyProduct({
        Product: this.product.Name,
        Price: this.product.Price,
      })
      .subscribe((data) => {
        console.log(data)
        if (data.success) {
          this.showBuy = true
        }
      });
  }

  deleteComment(comment: any) {
    console.log("Comment to delete");
    console.log(comment);

    this.productService.deleteComment(comment).subscribe((data) => {
      console.log("se ha eliminado el comentario");
      console.log(data);
    });
  }
}
