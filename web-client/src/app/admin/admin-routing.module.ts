import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { ListComponent as JobList } from './job/list/list.component';
import { DetailsComponent as JobDetails } from './job/details/details.component';
import { PreviewComponent as JobPreview } from './job/preview/preview.component';
import { ListComponent as LocationList } from './location/list/list.component';
import { DetailsComponent as LocationDetails } from './location/details/details.component';
import { ListComponent as PersonList } from './person/list/list.component';
import { DetailsComponent as PersonDetails } from './person/details/details.component';
import { ListComponent as ApplicationList } from './application/list/list.component';
import { DetailsComponent as ApplicationDetails } from './application/details/details.component';

const routes: Routes = [
  {
    path: '',
    component: DashboardComponent,
    children: [
      {
        path: 'applications/:page',
        component: ApplicationList,
        pathMatch: 'full',
      },
      {
        path: 'application',
        children: [
          {
            path: 'new',
            pathMatch: 'full',
            component: ApplicationDetails,
          },
          {
            path: 'details/:applicationId',
            pathMatch: 'full',
            component: ApplicationDetails,
          },
        ],
      },

      {
        path: 'locations/:page',
        component: LocationList,
        pathMatch: 'full',
      },
      {
        path: 'location',
        children: [
          {
            path: 'new',
            pathMatch: 'full',
            component: LocationDetails,
          },
          {
            path: 'details/:locationId',
            pathMatch: 'full',
            component: LocationDetails,
          },
        ],
      },

      {
        path: 'persons/:page',
        component: PersonList,
        pathMatch: 'full',
      },
      {
        path: 'person',
        children: [
          {
            path: 'person',
            pathMatch: 'full',
            component: PersonDetails,
          },
          {
            path: 'details/:personId',
            pathMatch: 'full',
            component: PersonDetails,
          },
        ],
      },

      {
        path: 'jobs/:page',
        component: JobList,
        pathMatch: 'full',
      },
      {
        path: 'job',
        children: [
          {
            path: 'preview/:jobId',
            pathMatch: 'full',
            component: JobPreview,
          },
          {
            path: 'new',
            pathMatch: 'full',
            component: JobDetails,
          },
          {
            path: 'details/:jobId',
            pathMatch: 'full',
            component: JobDetails,
          },
        ],
      },
    ],
  },
  { path: '**', redirectTo: '' },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AdminRoutingModule {}
