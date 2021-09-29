import { DetailComponentRouteDataInterface } from "../../interfaces/route-data/detail-component-route-data.interface";
import { DetailComponentModeEnum } from "../../enums/detail-component/detail-component-mode.enum";

export class DetailComponentModeModel {
  /**
   * creates a new object for the given route data
   * @param mode
   */
  static createRouteDataForDetailMode(
    mode: DetailComponentModeEnum
  ): DetailComponentRouteDataInterface {
    return <DetailComponentRouteDataInterface>{
      mode: mode,
    };
  }
}
