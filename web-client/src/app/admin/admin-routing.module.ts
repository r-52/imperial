import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { DashboardComponent } from "./dashboard/dashboard.component";
import { ApplicationListComponent } from "./application/application-list/application-list.component";
import { ApplicationDetailsComponent } from "./application/application-details/application-details.component";
import { LocationListComponent } from "./location/location-list/location-list.component";
import { LocationDetailsComponent } from "./location/location-details/location-details.component";
import { PersonListComponent } from "./person/person-list/person-list.component";
import { PersonDetailsComponent } from "./person/person-details/person-details.component";
import { JobListComponent } from "./job/job-list/job-list.component";
import { JobPreviewComponent } from "./job/job-preview/job-preview.component";
import { JobDetailsComponent } from "./job/job-details/job-details.component";
import { DetailComponentModeModel } from "../shared/models/detail-component/detail-component-mode.model";
import { DetailComponentModeEnum } from "../shared/enums/detail-component/detail-component-mode.enum";

const routes: Routes = [
  {
    path: "",
    component: DashboardComponent,
    children: [
      {
        path: "applications/:page",
        component: ApplicationListComponent,
        pathMatch: "full",
      },
      {
        path: "application",
        children: [
          {
            path: "new",
            pathMatch: "full",
            component: ApplicationDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.newMode
            ),
          },
          {
            path: "details/:applicationId",
            pathMatch: "full",
            component: ApplicationDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.readOnlyMode
            ),
          },
        ],
      },

      {
        path: "locations/:page",
        component: LocationListComponent,
        pathMatch: "full",
      },
      {
        path: "location",
        children: [
          {
            path: "new",
            pathMatch: "full",
            component: LocationDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.newMode
            ),
          },
          {
            path: "details/:locationId",
            pathMatch: "full",
            component: LocationDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.readOnlyMode
            ),
          },
          {
            path: "edit/:locationId",
            pathMatch: "full",
            component: LocationDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.editingMode
            ),
          },
        ],
      },

      {
        path: "persons/:page",
        component: PersonListComponent,
        pathMatch: "full",
      },
      {
        path: "person",
        children: [
          {
            path: "new",
            pathMatch: "full",
            component: PersonDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.newMode
            ),
          },
          {
            path: "details/:personId",
            pathMatch: "full",
            component: PersonDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.readOnlyMode
            ),
          },
          {
            path: "details/:personId",
            pathMatch: "full",
            component: PersonDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.editingMode
            ),
          },
        ],
      },

      {
        path: "jobs/:page",
        component: JobListComponent,
        pathMatch: "full",
      },
      {
        path: "job",
        children: [
          {
            path: "preview/:jobId",
            pathMatch: "full",
            component: JobPreviewComponent,
          },
          {
            path: "new",
            pathMatch: "full",
            component: JobDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.newMode
            ),
          },
          {
            path: "details/:jobId",
            pathMatch: "full",
            component: JobDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.readOnlyMode
            ),
          },
          {
            path: "edit/:jobId",
            pathMatch: "full",
            component: JobDetailsComponent,
            data: DetailComponentModeModel.createRouteDataForDetailMode(
              DetailComponentModeEnum.editingMode
            ),
          },
        ],
      },
    ],
  },
  { path: "**", redirectTo: "" },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AdminRoutingModule {}
