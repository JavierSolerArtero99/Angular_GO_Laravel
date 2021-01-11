import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PanelAdminRoutingModule } from './panel-admin-routing.module';
import { PanelComponent } from './panel/panel.component';

@NgModule({
  declarations: [PanelComponent],
  imports: [
    CommonModule,
    PanelAdminRoutingModule
  ]
})
export class PanelAdminModule { }
