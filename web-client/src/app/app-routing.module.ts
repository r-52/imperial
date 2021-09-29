import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

const routes: Routes = [
  {
    path: "",
    loadChildren: () =>
      import("./homepage/homepage.module").then((h) => h.HomepageModule),
  },
  {
    path: "admin",
    loadChildren: () =>
      import("./admin/admin.module").then((a) => a.AdminModule),
  },
  {
    path: "client",
    loadChildren: () =>
      import("./client/client.module").then((c) => c.ClientModule),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
