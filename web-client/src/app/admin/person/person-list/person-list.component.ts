import { Component } from "@angular/core";
import { BaseComponent } from "../../../shared/components/base/base/base.component";
import { Router } from "@angular/router";

@Component({
  selector: "app-person-list",
  templateUrl: "./person-list.component.html",
  styleUrls: ["./person-list.component.scss"],
})
export class PersonListComponent extends BaseComponent {
  constructor(private readonly _router: Router) {
    super();
  }

  /* **** Event listener **** */
  public onNewClick(): void {
    this._router.navigateByUrl("/admin/person/new");
  }
}
