import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
// import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {ProgressBarComponent} from './progress_bar/progress_bar.componenet';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatButtonModule} from '@angular/material/button';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {FontAwesomeModule} from '@fortawesome/angular-fontawesome';
import {MatInputModule} from '@angular/material/input';
import {HttpClientModule} from '@angular/common/http';
import {CompletedWorkComponent} from './completed_work/CompletedWorkComponent';
import {ApiService} from './http/ApiService';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {RefreshWorkDone} from './notification/RefreshWorkDone';
import {PromisePopupComponent} from './promise_popup/promise_popup.component';

@NgModule({
  declarations: [
    AppComponent,
    ProgressBarComponent,
    CompletedWorkComponent,
    PromisePopupComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    // BrowserAnimationsModule,
    MatToolbarModule,
    MatProgressSpinnerModule,
    MatButtonModule,
    BrowserAnimationsModule,
    FontAwesomeModule,
    MatInputModule,
    HttpClientModule,
    NgbModule,
  ],
  providers: [ApiService, RefreshWorkDone],
  bootstrap: [AppComponent]
})
export class AppModule {
}
