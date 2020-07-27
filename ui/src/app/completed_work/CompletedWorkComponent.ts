import {Component, OnInit} from '@angular/core';
import {ApiService} from '../http/ApiService';
import {TimeBySprintName} from '../poto/TimeBySprintName';
import {RefreshWorkDone} from '../notification/RefreshWorkDone';
import {EventTypes} from '../constants/EventTypes';

@Component({
  selector: 'app-completed-work',
  templateUrl: 'completed-work.component.html',
  styleUrls: ['./completed-work.component.styl'],
})
export class CompletedWorkComponent implements OnInit {
  completed: TimeBySprintName[];

  constructor(private api: ApiService, private refreshWorkDone: RefreshWorkDone) {
  }

  ngOnInit(): void {
    this.listSprints();
    this.refreshWorkDone.events$.forEach(event => {
      if (event === EventTypes.REFRESH_WORK_DONE) {
        this.listSprints();
      }
    });
  }

  private listSprints() {
    this.api.listSprints(sprints => {
      this.completed = sprints;
    });
  }
}
