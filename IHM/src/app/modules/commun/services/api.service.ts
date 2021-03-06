import {Injectable} from '@angular/core';
import {environment} from '../../../../environments/environment';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import {catchError, map} from 'rxjs/operators';



@Injectable({
    providedIn: 'root'
})
export class ApiService {
    /**
     * Url API Zone
     */

    /**
     * Airport
     */
    public static API_GET_ALL_AIRPORTS = 'airports';
    public static API_GET_AIRPORT_DATA = 'airport/';

    /**
     * Sensors
     */
    public static API_GET_SENSOR = 'sensor/';

    /**
     * Average
     */
    public static API_AVERAGE = 'average/';

    /**
     *
     * End: Url API Zone
     *
     */

    private apiUrl = environment.API_ENDPOINT;
    private readonly httpOptionsGet: HttpHeaders = new HttpHeaders();
    private token = '';

    constructor(private http: HttpClient) {
        if (environment.HTTP_OPTIONS) {
            if (environment.HTTP_OPTIONS_CORS) {
                this.httpOptionsGet = new HttpHeaders({
                    'Content-Type': 'application/json',
                });
            } else {
                this.httpOptionsGet = new HttpHeaders({
                    'Content-Type': 'application/json',

                });
            }
        }
    }

    doGetRequest(url: string, withAuthorization: boolean = true): Observable<object> {
        return this.http.get(this.apiUrl + url, {
            'headers': this.getHeaders(this.httpOptionsGet, withAuthorization)
        }).pipe(map((response: Response) => response), catchError(this.handleErrors('doGetRequest', '')));
    }

    handleErrors(operation = 'operation', result?: string) {
        return (error: any): Observable<object> => {
            if (error !== undefined) {
                result = error;
            } else {
                result = 'Sorry, an error occured';
            }
            return throwError(result);
        };
    }

    getHeaders(headers: HttpHeaders, withAuthentication: boolean): HttpHeaders {
        if (!withAuthentication) {
            return headers;
        }

        if ('' !== this.token && null !== this.token && undefined !== this.token) {
            headers.append('Authorization', 'Bearer ' + this.token);
        }

        return headers;
    }
}
