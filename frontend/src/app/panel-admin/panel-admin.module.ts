import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";

import { PanelAdminRoutingModule } from "./panel-admin-routing.module";
import { PanelComponent } from "./panel/panel.component";
import { SharedModule } from "../shared";

@NgModule({
  imports: [SharedModule, PanelAdminRoutingModule],
  declarations: [PanelComponent],
})
export class PanelAdminModule {}
