import {Component} from '@angular/core';
import {EventBus} from './notification/EventBus';
import {EventTypes} from './constants/EventTypes';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.styl']
})
export class AppComponent {
  constructor(private notifier: EventBus) {
  }

  title = 'ui';

  makePromiseNotify() {
    console.log('sending show event');
    this.notifier.newEvent(EventTypes.SHOW_PROMISE_POPIP);
  }
}
