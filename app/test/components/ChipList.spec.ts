
import { describe, it, expect, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import ChipList from '../../src/components/ChipList.vue'
import { wrapInArray } from 'vuetify/lib/util/helpers.mjs'

describe('ChipList', () => {
  
  let wrapper: any
  beforeEach(() => {
    wrapper = mount(ChipList, {
      props: {
        items: ["React", "Node", "Vue"],
        maxHeight: '200px'
      }
    })
  
  })

  it('should render your skills in label when label prop is not provided', () => {
    const label = wrapper.get('[data-testid="label"]')
    expect(label.text()).toBe('Yours Skills')
  })

  it('should render the component when label prop is provided', async () => {
    await wrapper.setProps({
      label: 'New Skills'
    })

    const label = wrapper.get('[data-testid="label"]')
    expect(label.text()).toBe('New Skills')
  })

  it('icon must be appear when list has one element',async () => {
    wrapper = mount(ChipList, {
      props: {
        items: [],
        maxHeight: '200px'
      }
    })

    let icon = wrapper.find('[data-testid="icon"]')
    expect(icon.exists()).toBe(false)

    await wrapper.setProps({
      items: ["React", "Node", "Vue"],
    })

    icon = wrapper.find('[data-testid="icon"]')
    expect(icon.exists()).toBe(true)
  })
})