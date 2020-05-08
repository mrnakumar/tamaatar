import {Component, Input} from '@angular/core';
import {interval, Subscription} from 'rxjs';

@Component({
  selector: 'app-progress-bar',
  templateUrl: 'progress_bar.component.html'
})
export class ProgressBarComponent {
  pbValue = 100;
  currSec = 0;

  startTimer(seconds: number) {
    this.currSec = seconds;
    const sub = interval(1000).subscribe((sec) => {
      this.pbValue = 100 - sec * 100 / seconds;
      this.currSec = sec;
      if (this.currSec === seconds) {
        sub.unsubscribe();
      }
    });
  }
}
