import {Injectable} from '@angular/core';
import {Subject} from 'rxjs';

@Injectable()
export class EventBus {
  private subject = new Subject<string>();

  newEvent(event) {
    this.subject.next(event);
  }

  get events$() {
    return this.subject.asObservable();
  }
}
