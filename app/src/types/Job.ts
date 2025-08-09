import type { Seniority } from "../enums/seniority";
import type { Modality } from "../enums/modality";
import type { Category } from "../enums/category";
import type { HostScrapper } from "./HostScrapper";

export interface Job {
    position: string;
    level: Seniority[];
    minimumSalary: number;
    maximumSalary: number;
    skills: string[];
    modalities: Modality[];
    description: string;
    company: string;
    location: string[];
    url: string;
    createdAt?: Date;
    web: string;
    host: number;
    tags:string[]
    contractType:string
    isRecentJob:boolean
    categories:Category[]
}


export interface JobRequest{
    location:string,
    level:Seniority,
    skills:string[],
    modalities:Modality,
    minimumSalaryExpectation:number,
    maximumSalaryExpectation:number,
    position:string,
    category:Category,
    hostSelected:number[]
}

export interface MatchAnalizer{
    porcentSalary: number,
    porcentSkills: number,
    porcentModalities: number,
    porcentLevels: number,
    porcentPosition: number,
    totalPorcent: number,
    skillMatches: string[],
    salaryMessage: string   
}

export interface JobScrapeated{
    job:Job
    matchAnalizer:MatchAnalizer
}