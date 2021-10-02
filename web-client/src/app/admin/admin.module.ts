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
import { NavigationListComponent } from "./shared/components/navigation/navigation-list/navigation-list.component";
import { PersonListTableComponent } from "./shared/components/person/person-list-table/person-list-table.component";
import { LocationListTableComponent } from "./shared/components/location/location-list-table/location-list-table.component";
import { ScheduleListTableComponent } from "./shared/components/schedule/schedule-list-table/schedule-list-table.component";
import { ScheduleListComponent } from "./schedule/schedule-list/schedule-list.component";
import { ScheduleDetailsComponent } from "./schedule/schedule-details/schedule-details.component";
import { SharedModule } from "../shared/shared.module";
import { ReactiveFormsModule } from "@angular/forms";
import { CreateFooterComponent } from "./shared/components/details/footer/create-footer/create-footer.component";
import { SaveFooterComponent } from "./shared/components/details/footer/save-footer/save-footer.component";
import { EditFooterComponent } from "./shared/components/details/footer/edit-footer/edit-footer.component";
import { BaseFooterComponent } from "./shared/components/details/footer/base-footer/base-footer.component";

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
    NavigationListComponent,
    PersonListTableComponent,
    LocationListTableComponent,
    ScheduleListTableComponent,
    ScheduleListComponent,
    ScheduleDetailsComponent,
    CreateFooterComponent,
    SaveFooterComponent,
    EditFooterComponent,
    BaseFooterComponent,
  ],
  imports: [
    CommonModule,
    AdminRoutingModule,
    MaterialModule,
    SharedModule,
    ReactiveFormsModule,
  ],
})
export class AdminModule {}
