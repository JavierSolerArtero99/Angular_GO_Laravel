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
  deleteComment(comment: any): any {
    console.log(comment);
    

    // return this.apiService
    //   .deleteGoComments(`/comment`, {
    //     UserID: comment.UserID,
    //     ProductID: comment.ProductID,
    //     Message: comment.Message,
    //   })
    //   .pipe(map((data) => data));

  }
}
