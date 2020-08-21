import {Component, OnInit} from '@angular/core';
import {PromisePoTo} from '../poto/PromisePoTo';
import {faPlus} from '@fortawesome/free-solid-svg-icons';
import {EventBus} from '../notification/EventBus';
import {EventTypes} from '../constants/EventTypes';

@Component({
  selector: 'app-promise-popup',
  templateUrl: 'promise-popup.component.html',
  styleUrls: ['./promise-popup.component.styl']
})
export class PromisePopupComponent implements OnInit {
  faPlus = faPlus;
  promises: PromisePoTo[] = [new PromisePoTo('', 0)];
  show = false;

  constructor(private notifier: EventBus) {
  }

  ngOnInit(): void {
    console.log('init');
    this.notifier.events$.subscribe(next => {
      console.log('out');
      if (next === EventTypes.SHOW_PROMISE_POPIP) {
        console.log('in');
        this.showPopup();
      }
    });
  }

  addPromise(name: string, duration: number) {
    this.promises.push(new PromisePoTo(name, duration));
  }

  showPopup() {
    console.log(this.promises.length + ' ovaova');
    this.show = true;
  }

  cancelPopup() {
    this.resetPopup();
    this.show = false;
  }

  private resetPopup() {
    this.promises = [];
  }
}
