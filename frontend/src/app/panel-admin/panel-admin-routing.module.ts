import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { RoleGuard } from "../core/services/role-guard.service";
import { PanelComponent } from "./panel/panel.component";

const routes: Routes = [
  {
    path: '',
    component: PanelComponent,
  },
];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class PanelAdminRoutingModule {}
