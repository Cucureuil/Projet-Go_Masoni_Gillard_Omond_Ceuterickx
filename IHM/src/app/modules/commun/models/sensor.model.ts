
export class SensorModel {
    private _idSensor: string;
    private _idAirport: string;
    private _value: number;
    private _dateMeasure: Date;


    get idSensor(): string {
        return this._idSensor;
    }

    set idSensor(value: string) {
        this._idSensor = value;
    }

    get idAirport(): string {
        return this._idAirport;
    }

    set idAirport(value: string) {
        this._idAirport = value;
    }

    get value(): number {
        return this._value;
    }

    set value(value: number) {
        this._value = value;
    }

    get dateMeasure(): Date {
        return new Date((this._dateMeasure as any) * 1000);
    }

    set dateMeasure(value: Date) {
        this._dateMeasure = value;
    }
}
