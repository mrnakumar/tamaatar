import {Component} from '@angular/core';
import {RefreshWorkDone} from './notification/RefreshWorkDone';
import {EventTypes} from './constants/EventTypes';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.styl']
})
export class AppComponent {
  constructor(private notifier: RefreshWorkDone) {
  }

  title = 'ui';

  makePromiseNotify() {
    console.log('sending show event');
    this.notifier.newEvent(EventTypes.SHOW_PROMISE_POPIP);
  }
}
