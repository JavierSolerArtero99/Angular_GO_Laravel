import { Injectable } from "@angular/core";
import { ApiService } from "../../core/services/api.service";
import { Observable } from "rxjs";
import { Product } from "./product.model";
import { HttpParams } from "@angular/common/http";
import { ProductListConfig } from "./product-list-config.model";
import { map } from "rxjs/operators";

@Injectable({
  providedIn: "root",
})
export class ProductService {
  constructor(private apiService: ApiService) {}

  getProductList(
    config: ProductListConfig
  ): Observable<{
    products: Product[];
    productsCount: number;
  }> {
    const params = {};

    return this.apiService.getGoProducts(
      "",
      new HttpParams({ fromObject: params })
    );
  }

  getSingleProduct(
    productName: string
  ): Observable<{
    product: Product;
  }> {
    const params = {};

    return this.apiService.getGoProducts(`/product?name=${productName}`);
  }

  postProduct(product: any): Observable<Product> {
    return this.apiService.post("/products", product);
  }

  postComment(comment: any): Observable<Comment> {
    console.log(comment);
    return this.apiService
      .postGoProducts(`/comment`, {
        UserID: comment.UserID,
        ProductID: comment.ProductID,
        Message: comment.Message,
      })
      .pipe(map((data) => data));
  }

  // deleteComment(comment: any): Observable<Comment> {
  deleteComment(comment: any): Observable<Comment> {
    return this.apiService
      .deleteGoComments(`/comment/${comment.ID}`)
      .pipe(map((data) => data));
  }

  likeProduct(product: any, userId: any): Observable<Comment> {
    return this.apiService
      .postGoProducts(`/like/${product.Name}/${userId}`)
      .pipe(map((data) => data));
  }

  unLike(product: Product, id: any) {
    console.log("Deleting MG")
    return null
  }

  buyProduct(product): any {
    return this.apiService
      .postGoProducts(`/buy`, product)
      .pipe(map((data) => data));
  }
}
