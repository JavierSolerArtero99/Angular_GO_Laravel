import { Component, Input, OnInit } from "@angular/core";

import { Product } from "./../shared/product.model";

@Component({
  selector: "app-product-preview",
  templateUrl: "./product-preview.component.html",
})
export class ProductPreviewComponent implements OnInit {
  @Input() product: Product;

  ngOnInit(): void {
    console.log(this.product)
}
}
