<template>
  <date-range-picker
    ref="picker"
    opens="center"
    :locale-data="{ firstDay: 1, format: 'DD-MM-YYYY HH:mm:ss' }"
    minDate="2019-05-02 04:00:00" maxDate="2020-12-26 14:00:00"
    :singleDatePicker="false"
    :timePicker="true"
    :timePicker24Hour="true"
    :showWeekNumbers="true"
    :showDropdowns="true"
    :autoApply="true"
    v-model="dateRange"
    @update="updateValues"
    @toggle="checkOpen"
    :dateFormat="dateFormat"
  >
    <template v-slot:input="picker" style="min-width: 350px;">
        {{ picker.startDate }} - {{ picker.endDate }}
    </template>
  </date-range-picker>
</template>
<script>
import DateRangePicker from 'vue2-daterange-picker'
import 'vue2-daterange-picker/dist/vue2-daterange-picker.css'

export default {
  components: {
    DateRangePicker
  },
  data () {
    return {
      dateRange: {
        startDate: '2019-05-02 04:00:00',
        endDate: '2020-12-26 14:00:00'
      }
    }
  },
  methods: {
    updateValues () {
    },
    checkOpen () {
    },
    dateFormat () {
      return new Date()
    }
  }
}
</script>