import {Component} from '@angular/core';
import {PromisePoTo} from '../poto/PromisePoTo';
import {faPlus} from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-promise-popup',
  templateUrl: 'promise-popup.component.html',
  styleUrls: ['./promise-popup.component.styl']
})
export class PromisePopupComponent {
  faPlus = faPlus;
  promises: PromisePoTo[] = [new PromisePoTo('', 0)];
  show = false;

  addPromise(name: string, duration: number) {
    this.promises.push(new PromisePoTo(name, duration));
  }

  showPopup() {
    this.resetPopup();
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
