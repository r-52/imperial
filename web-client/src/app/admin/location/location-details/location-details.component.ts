import { Component, OnInit } from "@angular/core";
import { BaseComponent } from "../../../shared/components/base/base/base.component";
import { ActivatedRoute } from "@angular/router";
import { combineLatest } from "rxjs";
import { TitleGeneratorService } from "../../../shared/services/title/title-generator.service";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";

import { Location } from "@angular/common";
import { BaseDetailComponent } from "../../../shared/components/base/base-detail/base-detail.component";

@Component({
  selector: "app-location-details",
  templateUrl: "./location-details.component.html",
  styleUrls: ["./location-details.component.scss"],
})
export class LocationDetailsComponent
  extends BaseDetailComponent
  implements OnInit
{
  /**
   * the form
   */
  public form: FormGroup = this._formBuilder.group({
    name: ["", Validators.required],
    street1: ["", Validators.required],
    street2: [""],
    zip: ["", Validators.required],
    city: ["", Validators.required],
    country: [""],
  });

  constructor(
    private readonly _route: ActivatedRoute,
    private readonly _titleGenerator: TitleGeneratorService,
    private readonly _formBuilder: FormBuilder,
    private readonly _location: Location
  ) {
    super();
  }

  public ngOnInit() {
    this._loadSubscriptions();
  }

  private _loadSubscriptions(): void {
    this._subscriptions.push(
      combineLatest([this._route.data]).subscribe((next) => {
        const routeData = next[0];
        this.detailMode = routeData.mode;

        this.title = this._titleGenerator.makeTitle(
          routeData.mode,
          "",
          "Location"
        );
      })
    );
  }

  /* **** Event listener **** */
  public onCancelClick(): void {
    this._location.back();
  }

  public onCreateClick(): void {}

  public onSaveClick(): void {}

  public onEditClick(): void {}

  /* **** getter **** */
  public get name() {
    return this.form.get("name");
  }

  public get street1() {
    return this.form.get("street1");
  }

  public get street2() {
    return this.form.get("street2");
  }

  public get zip() {
    return this.form.get("zip");
  }

  public get city() {
    return this.form.get("city");
  }

  public get country() {
    return this.form.get("country");
  }
}
