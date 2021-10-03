import { Component, OnInit } from "@angular/core";
import { Location } from "@angular/common";
import { BaseComponent } from "../../../shared/components/base/base/base.component";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { TitleGeneratorService } from "../../../shared/services/title/title-generator.service";
import { BaseDetailComponent } from "../../../shared/components/base/base-detail/base-detail.component";

@Component({
  selector: "app-person-details",
  templateUrl: "./person-details.component.html",
  styleUrls: ["./person-details.component.scss"],
})
export class PersonDetailsComponent
  extends BaseDetailComponent
  implements OnInit
{
  public form: FormGroup = this._formBuilder.group({
    firstName: ["", Validators.required],
    lastName: ["", Validators.required],
    email: ["", Validators.required, Validators.email],
    phone: [""],
    position: [""],
    isActive: ["", Validators.required],
  });

  public passwordForm: FormGroup = this._formBuilder.group({
    password: ["", Validators.required],
    confirmPassword: ["", Validators.required],
  });

  constructor(
    private readonly _formBuilder: FormBuilder,
    private readonly _location: Location,
    private readonly _titleGenerator: TitleGeneratorService
  ) {
    super();
  }

  /* **** life cycle **** */

  public ngOnInit(): void {
    this._titleGenerator.makeTitle(this.detailMode, "", "Person");
  }

  /* **** Event listener **** */
  public onCancelClick(): void {
    this._location.back();
  }

  public onCreateClick(): void {}

  public onSaveClick(): void {}

  public onEditClick(): void {}

  /* **** Getter **** */
  public get password() {
    return this.passwordForm.get("password");
  }

  public get confirmPassword() {
    return this.passwordForm.get("confirmPassword");
  }

  public get firstName() {
    return this.form.get("firstName");
  }

  public get lastName() {
    return this.form.get("lastName");
  }

  public get position() {
    return this.form.get("position");
  }

  public get email() {
    return this.form.get("email");
  }

  public get phone() {
    return this.form.get("phone");
  }

  public get isActive() {
    return this.form.get("isActive");
  }
}
