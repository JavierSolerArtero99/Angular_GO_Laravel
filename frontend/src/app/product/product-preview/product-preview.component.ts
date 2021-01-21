import { Component, Input, OnInit } from "@angular/core";
import { ProductService } from "../shared/product.service";

import { Product } from "./../shared/product.model";

@Component({
  selector: "app-product-preview",
  templateUrl: "./product-preview.component.html",
})
export class ProductPreviewComponent implements OnInit {
  @Input() product: Product;

  constructor(private productsService: ProductService) {

  }

  ngOnInit(): void {
    console.log(this.product);
  }

  likeProduct() {
    // dandole like al producto
    this.productsService.likeProduct(this.product).subscribe((data) => {
      this.product.likes++
    });
  }
}
