import { Component, Input, OnInit } from "@angular/core";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";
import { StatsService } from "../../core/services/stats.service";

@Component({
  selector: "app-panel",
  templateUrl: "./panel.component.html",
  styles: [],
})
export class PanelComponent implements OnInit {
  bestsBuys: any = [];
  totalAmount: number = 0;
  loadingProducts: boolean = true;
  currentUsers: number = 0;
  moneyEarned: number = 0;
  totalUsers: number = 0;
  productForm: FormGroup;
  commentMessage = new FormControl();
  productName = new FormControl();
  productPrice = new FormControl();
  productImage = new FormControl();
  productDescription = new FormControl();

  constructor(private fb: FormBuilder, private statsService: StatsService) {}

  isNewProduct: boolean = false;

  ngOnInit() {
    this.statsService.getCurrentUsersCache().subscribe((data) => {
      this.currentUsers = data.current_users;
    });

    this.statsService.getValoredProducts().subscribe((data) => {
      data.buys.sort(function (a, b) {
        if (a.TimesBuyed < b.TimesBuyed) {
          return 1;
        }
        if (a.TimesBuyed > b.TimesBuyed) {
          return -1;
        }
        return 0;
      });

      let cont = 0;
      data.buys.forEach(buy => {
        this.totalAmount += buy.Price * buy.TimesBuyed
        if (cont < 3) this.bestsBuys.push(buy);
      });

      this.loadingProducts = false;
    });

    this.productForm = this.fb.group({
      productName: "",
      productPrice: "",
      productImage: "",
      productDescription: "",
    });
  }

  toogleNewProductForm(): void {
    this.isNewProduct = !this.isNewProduct;
  }

  submitProduct() {
    console.log(this.productForm.getRawValue());
  }
}
