import {Component, OnDestroy, OnInit} from '@angular/core';
import {interval, Subscription} from 'rxjs';
import {faPlus} from '@fortawesome/free-solid-svg-icons';
import {ApiService} from '../http/ApiService';
import {RefreshWorkDone} from '../notification/RefreshWorkDone';
import {EventTypes} from '../constants/EventTypes';

@Component({
  selector: 'app-progress-bar',
  templateUrl: 'progress_bar.component.html',
  styleUrls: ['./progress_bar.component.styl']
})
export class ProgressBarComponent implements OnInit, OnDestroy {
  faPlus = faPlus;
  pbValue: number;
  currSec = 0;
  sprintName: string;
  durationInSeconds: number;
  remainingDurationInSeconds = 10 * 60;
  isRunning = false;
  subscription: Subscription = null;
  alarm: Subscription = null;
  durationDiv: DurationDiv = null;
  context = new window.AudioContext();

  durations: DurationDiv[] = [
    {duration: 1, selected: true},
    {duration: 13, selected: false},
    {duration: 16, selected: false},
    {duration: 19, selected: false},
    {duration: 22, selected: false},
    {duration: 25, selected: false},
    {duration: 28, selected: false},
    {duration: 31, selected: false},
    {duration: 34, selected: false},
    {duration: 37, selected: false}
  ];

  constructor(private api: ApiService, private notifier: RefreshWorkDone) {
  }

  ngOnInit(): void {
    if (this.durationDiv === null) {
      this.durationDiv = this.durations[0];
    }
    this.setDuration(this.durationDiv);
    this.pbValue = 100;
  }

  ngOnDestroy(): void {
    this.context.close().then(__ => console.log('audio context stopped'));
  }

  startOrStopTimer() {
    if (!this.isRunning) {
      if (this.durationInSeconds > 0) {
        this.mayBeStartTimer();
      }
    } else {
      this.mayBeStopTimer();
      this.stopAlarm();
    }
  }

  private mayBeStartTimer() {
    this.currSec = this.durationInSeconds;
    this.isRunning = true;
    this.subscription = interval(1000).subscribe((sec) => {
      this.pbValue = 100 - sec * 100 / this.durationInSeconds;
      this.currSec = sec;
      this.remainingDurationInSeconds = this.durationInSeconds - sec;
      if (this.currSec === this.durationInSeconds) {
        this.subscription.unsubscribe();
        this.subscription = null;
        this.startAlarm();
      }
    });
  }

  private mayBeStopTimer() {
    this.recordSprint();
    if (this.subscription !== null) {
      this.subscription.unsubscribe();
      this.subscription = null;
    }
    this.isRunning = false;
    this.ngOnInit();
  }

  setDuration(durationDiv: DurationDiv) {
    if (!this.isRunning) {
      this.durationDiv = durationDiv;
      this.durationInSeconds = durationDiv.duration * 60;
      this.remainingDurationInSeconds = this.durationInSeconds;
      durationDiv.selected = true;
      this.durations.filter(duration => duration.duration !== durationDiv.duration)
        .forEach(duration => duration.selected = false);
    }
  }

  setSprintName(sprintName: string) {
    this.sprintName = sprintName;
  }

  getSprintName(): string {
    if (this.sprintName === undefined || this.sprintName.length === 0) {
      return 'Anonymous';
    }
    return this.sprintName;
  }

  formatRemainingDuration(): string {
    const copy = this.remainingDurationInSeconds;
    const minutes = Math.trunc(copy / 60);
    const minutesStr = minutes < 10 ? '0' + minutes : minutes + '';
    const seconds = copy % 60;
    const secondsStr = seconds < 10 ? '0' + seconds : seconds + '';
    return minutesStr + ':' + secondsStr;
  }

  private startAlarm() {
    this.alarm = interval(1000).subscribe((__) => {
      const oscillator = this.context.createOscillator();
      oscillator.connect(this.context.destination);
      const when = this.context.currentTime;
      oscillator.start(when);
      oscillator.stop(when + 0.1);
    });
  }

  private stopAlarm() {
    if (this.alarm !== null) {
      this.alarm.unsubscribe();
      this.alarm = null;
    }
  }

  private recordSprint() {
    const inMinutes = this.durationInSeconds / 60;
    const remaining = this.remainingDurationInSeconds / 60;
    const timeSpent = inMinutes - remaining;
    if (timeSpent > 0) {
      this.api.createSprint(this.sprintName, timeSpent, () => {
        this.notifier.newEvent(EventTypes.REFRESH_WORK_DONE);
      });
    }
  }
}

class DurationDiv {
  duration: number;
  selected: boolean;
}
