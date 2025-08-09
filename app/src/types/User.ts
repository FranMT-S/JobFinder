import { Modality } from "../enums/modality"

export interface User{
  description: string
  level: string
  skills: string[]
  location: string
  modality: Modality
  minSalary: number
  maxSalary: number
  category: number
  hostSelected: number[]
  maxPage: number
}