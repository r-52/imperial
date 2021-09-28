import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AdminRoutingModule } from './admin-routing.module';
import { DashboardComponent } from './dashboard/dashboard.component';
import { MaterialModule } from '../material.module';
import { ListComponent } from './job/list/list.component';
import { DetailsComponent } from './job/details/details.component';
import { PreviewComponent } from './job/preview/preview.component';

@NgModule({
  declarations: [DashboardComponent, ListComponent, DetailsComponent, PreviewComponent],
  imports: [CommonModule, AdminRoutingModule, MaterialModule],
})
export class AdminModule {}
