import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'weekday'
})
export class WeekdayPipe implements PipeTransform {


  transform(value: number): string {
    let weekday = new Map([
      [1, "星期一"],
      [2, "星期二"],
      [3, "星期三"],
      [4, "星期四"],
      [5, "星期五"],
      [6, "星期六"],
      [7, "星期日"]
    ]);
    return weekday.get(value);
  }

}
