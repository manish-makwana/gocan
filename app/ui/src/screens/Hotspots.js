import {useEffect, useState} from "react";
import axios from "axios";
import {CirclePacking} from "../components/CirclePacking";
import {Spinner} from "../components/Spinner";
import {DateSelector} from "../components/DateSelector";
import {Button} from 'primereact/button';

export function Hotspots({sceneId, appId, date}) {
    const [dateRange, setDateRange] = useState(date);
    const [analyze, setAnalyze] = useState(true);
    const [hospots, setHotspots] = useState();
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        let subscribed = true;
        if (analyze) {
            setLoading(true);
            let endpoint;
            if (appId) {
                endpoint = `/api/scenes/${sceneId}/apps/${appId}/hotspots`;
            } else {
                endpoint = `/api/scenes/${sceneId}/hotspots`;
            }
            let params = new URLSearchParams();
            if (dateRange.min) {
                params.append("after", dateRange.min);
            }
            if (dateRange.max) {
                params.append("before", dateRange.max);
            }
            axios.get(`${endpoint}?${params}`)
                .then(it => it.data)
                .then(it => {
                    if (subscribed) {
                        setHotspots(it);
                    }
                }).finally(() => {
                setLoading(false);
                setAnalyze(false);
            });
        }

        return () => subscribed = false;
    }, [sceneId, appId, analyze, dateRange]);

    let screen;

    if (loading) {
        screen = <Spinner/>;
    } else if (hospots) {
        screen = <div style={{display: "flex", justifyContent: "center"}}>
            <CirclePacking width={975} height={975} data={hospots}/>
        </div>;
    } else {
        screen = <><p>Unable to get hotspots</p></>
    }

    return <>
        <div className="card mt-4">
            <div className="flex align-items-center">
                <DateSelector min={date.min} max={date.max} onChange={e => setDateRange(e)}/>
                <Button label="Submit" onClick={e => setAnalyze(true)}/>
            </div>
        </div>
        {screen}
    </>
}