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
  products: Product[];
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

    // obteniendo los datos del producto
    this.productService.getSingleProduct(productName).subscribe((data) => {
      this.product = data.product;
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

  deleteComment(comment: any) {
    console.log("Comentario a borrar");
    console.log(comment);

    this.productService.deleteComment(comment)
    
    // this.productService.deleteComment(comment).subscribe((data) => {
    //   console.log("se ha hecho")
    //   console.log(data);
    // });
  }
}
