import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AdminRoutingModule } from './admin-routing.module';
import { DashboardComponent } from './dashboard/dashboard.component';
import { MaterialModule } from '../material.module';
import { ListComponent } from './job/list/list.component';
import { DetailsComponent } from './job/details/details.component';
import { PreviewComponent } from './job/preview/preview.component';
import { ApplicationDetailsComponent } from './application/application-details/application-details.component';
import { ApplicationListComponent } from './application/application-list/application-list.component';
import { JobListComponent } from './job/job-list/job-list.component';
import { JobDetailsComponent } from './job/job-details/job-details.component';
import { JobPreviewComponent } from './job/job-preview/job-preview.component';
import { LocationDetailsComponent } from './location/location-details/location-details.component';
import { LocationListComponent } from './location/location-list/location-list.component';
import { PersonListComponent } from './person/person-list/person-list.component';
import { PersonDetailsComponent } from './person/person-details/person-details.component';

@NgModule({
  declarations: [DashboardComponent, ListComponent, DetailsComponent, PreviewComponent, ApplicationDetailsComponent, ApplicationListComponent, JobListComponent, JobDetailsComponent, JobPreviewComponent, LocationDetailsComponent, LocationListComponent, PersonListComponent, PersonDetailsComponent],
  imports: [CommonModule, AdminRoutingModule, MaterialModule],
})
export class AdminModule {}
