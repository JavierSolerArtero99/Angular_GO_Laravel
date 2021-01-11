import { Injectable } from "@angular/core";
import { ApiService } from "../../core/services/api.service";
import { Observable } from "rxjs";
import { Product } from "./product.model";
import { HttpParams } from "@angular/common/http";
import { ProductListConfig } from "./product-list-config.model";

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
    // const params = {};

    // Object.keys(config.filters).forEach((key) => {
    //   params[key] = config.filters[key];
    // });

    // return this.apiService.get(
    //   "/products" + (config.type === "feed" ? "/feed" : ""),
    //   new HttpParams({ fromObject: params })
    // );

    const params = {};

    return this.apiService.getGoProducts(
      "",
      new HttpParams({ fromObject: params })
    );
  }
}
