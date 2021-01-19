import { Component, OnInit } from "@angular/core";

@Component({
  selector: "app-panel",
  templateUrl: "./panel.component.html",
  styles: [],
})
export class PanelComponent implements OnInit {
  constructor() {}

  isNewProduct: boolean = false;

  ngOnInit() {}

  toogleNewProductForm():void {
    this.isNewProduct = !this.isNewProduct;  
  }
}
