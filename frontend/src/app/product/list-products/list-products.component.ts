import { Component, Input, OnInit } from "@angular/core";
import { ProductListConfig } from "../shared/product-list-config.model";
import { Product } from "../shared/product.model";
import { ProductService } from "../shared/product.service";

@Component({
  selector: "app-list-products",
  templateUrl: "./list-products.component.html",
  styles: [],
})
export class ListProductsComponent implements OnInit {
  constructor(private productService: ProductService) {}

  @Input()
  set config(config: any) {
    if (config) {
      this.query = config;
      this.currentPage = 1;
      this.runQuery();
    }
  }

  loading = false;
  query: ProductListConfig;
  currentPage: number;
  results: Product[]

  ngOnInit(): void {
    this.loading = true;
    this.runQuery();
  }

  runQuery() {
    this.productService.getProductList(this.query).subscribe((data) => {
      this.loading = false;
      this.results = data.products
    });
  }
}
