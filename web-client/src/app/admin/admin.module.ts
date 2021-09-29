import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";

import { AdminRoutingModule } from "./admin-routing.module";
import { DashboardComponent } from "./dashboard/dashboard.component";
import { MaterialModule } from "../material.module";
import { ApplicationDetailsComponent } from "./application/application-details/application-details.component";
import { ApplicationListComponent } from "./application/application-list/application-list.component";
import { JobListComponent } from "./job/job-list/job-list.component";
import { JobDetailsComponent } from "./job/job-details/job-details.component";
import { JobPreviewComponent } from "./job/job-preview/job-preview.component";
import { LocationDetailsComponent } from "./location/location-details/location-details.component";
import { LocationListComponent } from "./location/location-list/location-list.component";
import { PersonListComponent } from "./person/person-list/person-list.component";
import { PersonDetailsComponent } from "./person/person-details/person-details.component";
import { AuthForgottenComponent } from "./auth/auth-forgotten/auth-forgotten.component";
import { AuthLoginComponent } from "./auth/auth-login/auth-login.component";
import { AuthRegisterComponent } from "./auth/auth-register/auth-register.component";

@NgModule({
  declarations: [
    DashboardComponent,
    ApplicationDetailsComponent,
    ApplicationListComponent,
    JobListComponent,
    JobDetailsComponent,
    JobPreviewComponent,
    LocationDetailsComponent,
    LocationListComponent,
    PersonListComponent,
    PersonDetailsComponent,
    AuthForgottenComponent,
    AuthLoginComponent,
    AuthRegisterComponent,
  ],
  imports: [CommonModule, AdminRoutingModule, MaterialModule],
})
export class AdminModule {}
